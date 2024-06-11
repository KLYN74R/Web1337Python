from nacl.encoding import HexEncoder
from nacl.signing import SigningKey
import requests
import json

# Import classes for other Klyntar API (see ./api directory)

from api.block_api import Web1337BlockApi
from api.misc_api import Web1337MiscApi
from api.state_api import Web1337StateApi
from api.txs_api import Web1337TxsApi








class Web1337(Web1337BlockApi,Web1337MiscApi,Web1337StateApi,Web1337TxsApi):

    def __init__(self, symbiotic_chain_id, workflow_version, node_url, proxy_url=None):
        self.symbiotic_chains = {}
        self.current_symbiotic_chain = symbiotic_chain_id
        self.symbiotic_chains[symbiotic_chain_id] = {
            "node_url": node_url,
            "workflow_version": workflow_version
        }

        if proxy_url:
            self.proxy = {
                "http": proxy_url,
                "https": proxy_url
            }
        else:
            self.proxy = None

    def get_request(self, url):
        node_url = self.symbiotic_chains[self.current_symbiotic_chain]["node_url"]
        response = requests.get(node_url + url, proxies=self.proxy)
        response.raise_for_status()
        return response.json()

    def post_request(self, url, payload):
        node_url = self.symbiotic_chains[self.current_symbiotic_chain]["node_url"]
        response = requests.post(node_url + url, json=payload, proxies=self.proxy)
        response.raise_for_status()
        return response.json()
    
    

    def get_transaction_template(self, workflow_version, creator, tx_type, nonce, fee, payload):
        return {
            "v": workflow_version,
            "creator": creator,
            "type": tx_type,
            "nonce": nonce,
            "fee": fee,
            "payload": payload,
            "sig": ""
        }

    def create_default_transaction(self, origin_shard, your_address, your_private_key, nonce, recipient, fee, amount_in_kly, rev_t=None):
        workflow_version = self.symbiotes[self.current_symbiote]["workflow_version"]
        payload = {
            "type": "D",
            "to": recipient,
            "amount": amount_in_kly
        }
        if rev_t is not None:
            payload["rev_t"] = rev_t

        transaction = self.get_transaction_template(workflow_version, your_address, "TX", nonce, fee, payload)
        signing_key = SigningKey(your_private_key, encoder=HexEncoder)
        message = f"{self.current_symbiote}{workflow_version}{origin_shard}TX{json.dumps(payload)}{nonce}{fee}".encode()
        transaction["sig"] = signing_key.sign(message).signature.hex()

        return transaction