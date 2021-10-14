
int CCoreWorkerRunTaskExecutionLoop(
    const char *node_ip_address, 
    int node_manager_port, 
    int num_workers, 
    const char *redis_ip, 
    const char *redis_password, 
    int redis_port,
    const char *raylet_ip_address,
    int local_mode);
