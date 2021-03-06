const ApiHost = import.meta.env.PROD ? import.meta.env.QOVERY_ROUTER_IMFK_URL || location.origin : 'https://imfk-alc972agc01og4tf-gtw.qovery.io' || 'http://localhost:1447'

export const urls = {
  settings: `${ApiHost}/api/v1/settings`,
  eval: `${ApiHost}/api/v1/mf/:type/eval`,
  plot: `${ApiHost}/api/v1/mf/:type/plot`,
  operationPlot: `${ApiHost}/api/v1/mf/operation/:operation/:type/plot`,
  operationEval: `${ApiHost}/api/v1/mf/operation/:operation/:type/eval`
}
