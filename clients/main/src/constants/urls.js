const ApiHost = import.meta.env.APP_API_HOST || 'http://localhost:1447'

export const urls = {
  settings: `${ApiHost}/api/v1/settings`,
  eval: `${ApiHost}/api/v1/mf/:type/eval`,
  plot: `${ApiHost}/api/v1/mf/:type/plot`
}
