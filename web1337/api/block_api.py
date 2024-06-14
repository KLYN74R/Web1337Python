
class Web1337BlockApi:    

    def get_block_by_id(self, block_id):
        return self.get_request("/block/" + block_id)

    def get_block_by_sid(self, shard, index):
        return self.get_request(f"/block_by_sid/{shard}/{index}")

    def get_latest_n_blocks_on_shard(self, shard, start, limit):
        return self.get_request(f"/latest_n_blocks/{shard}/{start}/{limit}")
