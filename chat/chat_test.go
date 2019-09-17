package chat

import "testing"

func TestAddAndRemoveUserToken(t *testing.T) {
	if err := AddUserToken("a2JVMzYyVUdYV01PVy82V0N1RmMvNlY1c29YUVoxNHhPVktUdTVid2VpU3RwUXpXTGo1b3JNSStKeHNJNFI4RmVQQ01nKzE5ZExNWFFGaDc4OUh3L3c1ay85dTVSNFdWRElDc01HaWRQMkx5"); err != nil {
		t.Error(err)
	}
	if err := RemoveUserToken("a2JVMzYyVUdYV01PVy82V0N1RmMvNlY1c29YUVoxNHhPVktUdTVid2VpU3RwUXpXTGo1b3JNSStKeHNJNFI4RmVQQ01nKzE5ZExNWFFGaDc4OUh3L3c1ay85dTVSNFdWRElDc01HaWRQMkx5"); err != nil {
		t.Error(err)
	}
}
