
class Web1337BlockApi:    

    def get_block_by_block_id(self, block_id):
        return self.get_request("/block/" + block_id)

    def get_block_by_sid(self, shard, sid):
        return self.get_request(f"/block_by_sid/{shard}/{sid}")
