const ApiHost = import.meta.env.PROD ? import.meta.env.QOVERY_ROUTER_IMFK_URL || location.origin : 'http://localhost:1447'

export const urls = {
  settings: `${ApiHost}/api/v1/settings`,
  eval: `${ApiHost}/api/v1/mf/:type/eval`,
  plot: `${ApiHost}/api/v1/mf/:type/plot`
}
