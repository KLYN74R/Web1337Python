class Web1337StateApi:
        
    def get_data_from_state(self, shard, cell_id):
        return self.get_request(f"/state/{shard}/{cell_id}")
            
    def get_pool_stats(self, pool_id):
        return self.get_request(f"/pool_stats/{pool_id}")