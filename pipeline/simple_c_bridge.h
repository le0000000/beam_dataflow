#ifndef THIRD_PARTY_PRIVACY_SANDBOX_AGGREGATION_PIPELINE_SIMPLE_C_BRIDGE_H_
#define THIRD_PARTY_PRIVACY_SANDBOX_AGGREGATION_PIPELINE_SIMPLE_C_BRIDGE_H_

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

const int const_value = 3;

int CGetValue() {
  return const_value;
}

#ifdef __cplusplus
}
#endif

#endif  // THIRD_PARTY_PRIVACY_SANDBOX_AGGREGATION_PIPELINE_DISTRIBUTED_POINT_FUNCTION_C_BRIDGE_H_
