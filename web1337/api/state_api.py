class Web1337StateApi:
        
    def get_from_state(self, shard, cell_id):
        return self.get_request(f"/state/{shard}/{cell_id}")
    
    def get_sync_state(self):
        return self.get_request("/sync_state")
    
    def get_symbiote_info(self):
        return self.get_request("/symbiote_info")

    def get_current_checkpoint(self):
        return self.get_request("/quorum_thread_checkpoint")


