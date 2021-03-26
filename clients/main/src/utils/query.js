export const qs = {
  get params () {
    return new URLSearchParams(location.search)
  }
}
