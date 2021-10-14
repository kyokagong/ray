#include "make_unique.hpp"
#include <msgpack.hpp>
#include "ray/common/id.h"
#include "ray/core_worker/context.h"
#include "ray/core_worker/store_provider/memory_store/memory_store.h"
#include "ray/core_worker/core_worker.h"


ray::core::CoreWorkerOptions GetCoreWorkerOptions();
ray::core::CoreWorker &GetCoreWorkerByCoreWorkerProcess();

bool NativePutRaw(const char *data);

const char *NavtiveGetRaw();
