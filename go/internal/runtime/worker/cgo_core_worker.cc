#include "ray/core_worker/core_worker.h"
#include "src/ray/protobuf/common.pb.h"

extern "C" {
    #include "cgo_core_worker.h"
}


int CCoreWorkerRunTaskExecutionLoop(
        const char *node_ip_address, 
        int node_manager_port, 
        int num_workers, 
        const char *redis_ip, 
        const char *redis_password, 
        int redis_port,
        const char *raylet_ip_address,
        int local_mode) {

    std::string session_dir = "/tmp/ray/session_dir";

    std::string store_socket = session_dir + "/sockets/plasma_store";

    std::string raylet_socket = session_dir + "/sockets/raylet";

    std::string log_dir = session_dir + "/logs";
    
    ray::gcs::GcsClientOptions gcs_options = ray::gcs::GcsClientOptions(std::string(redis_ip), redis_port, std::string(redis_password));

    ray::CoreWorkerOptions options;

    bool is_local_mode;
    if (local_mode > 0) {
        is_local_mode = true;
    } else {
        is_local_mode = false;
    }
    options.worker_type = ray::WorkerType::WORKER;
    options.language = Language::CPP;
    options.store_socket = store_socket;
    options.raylet_socket = raylet_socket;
    options.gcs_options = gcs_options;
    options.enable_logging = true;
    options.is_local_mode = is_local_mode;
    options.log_dir = log_dir;
    options.install_failure_signal_handler = true;
    options.node_ip_address = std::string(node_ip_address);
    options.node_manager_port = node_manager_port;
    options.raylet_ip_address = std::string(raylet_ip_address);
    options.driver_name = "go_worker";
    options.ref_counting_enabled = true;
    options.num_workers = num_workers;
    options.metrics_agent_port = -1;
    // options.task_execution_callback = callback;
    
    ray::CoreWorkerProcess::Initialize(options);
    ray::CoreWorkerProcess::RunTaskExecutionLoop();
    return 0;
}
