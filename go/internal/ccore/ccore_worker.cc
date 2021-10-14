#include "ray/common/id.h"
#include "ray/core_worker/context.h"
#include "ray/core_worker/store_provider/memory_store/memory_store.h"
#include "ccore_worker.h"


ray::core::CoreWorkerOptions GetCoreWorkerOptions() {
    ray::core::CoreWorkerOptions options;
    return options;
}

ray::core::CoreWorker &GetCoreWorkerByCoreWorkerProcess() {
    return ray::core::CoreWorkerProcess::GetCoreWorker();
}
