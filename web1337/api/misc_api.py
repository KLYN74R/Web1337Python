class Web1337MiscApi:
    def get_aggregated_finalization_proof_for_block(self, block_id):
        return self.get_request("/aggregated_finalization_proof/" + block_id)

    def get_general_info_about_kly_infrastructure(self):
        return self.get_request("/my_info")