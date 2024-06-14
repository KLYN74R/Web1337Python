
class Web1337EpochApi:    

    def get_current_epoch_on_thread(self, thread_id):
        return self.get_request(f"/current_epoch/{thread_id}")

    def get_current_leaders_on_shards(self):
        return self.get_request("/current_shards_leaders")

    def get_epoch_data_by_epoch_index(self, epoch_index):
        return self.get_request(f"/epoch_by_index/{epoch_index}")
    
    def get_aggregated_epoch_finalization_proof(self, epoch_index,shard):
        return self.get_request(f"/aggregated_epoch_finalization_proof/{epoch_index}/{shard}")
    
    def get_epoch_data_by_epoch_index(self, block_id):
        return self.get_request(f"/aggregated_finalization_proof/{block_id}")