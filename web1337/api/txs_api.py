class Web1337TxsApi:
    
    def get_transaction_receipt_by_id(self, tx_id):
        return self.get_request("/tx_receipt/" + tx_id)
    
    def send_transaction(self, transaction):
        return self.post_request("/transaction", transaction)

