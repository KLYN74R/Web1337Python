class Web1337MiscApi:
    
    def get_target_node_infrastructure_info(self):
        return self.get_request("/infrastructure_info")

    def get_chain_data(self):
        return self.get_request("/chain_info")
    
    def get_kly_evm_metadata(self):
        return self.get_request("/kly_evm_metadata")
    
    def get_synchronization_stats(self):
        return self.get_request("/synchronization_stats")