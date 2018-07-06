ETCDCTL_API=3 etcdctl put /sfc-controller/v2/config/network-node/vswitch2 '{"metadata":{"name":"vswitch2","labels":{"bar":"foo"}},"spec":{"node_type":"host","interfaces":[{"name":"GigabitEthernet2/0/1","if_type":"ethernet","ip_addresses":["192.168.16.1/24"],"parent":"vswitch2"},{"name":"vxlan_endpoint","if_type":"vxlan_tunnel","ip_addresses":["200.200.200.1/24"],"parent":"vswitch2"}]},"status":{"msg":["hello status"]}}'
