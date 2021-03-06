export const defaultParams = {
  trimf: [2 / 10, 5 / 10, 8 / 10],
  trapmf: [2 / 10, 4 / 10, 6 / 10, 7.5 / 10],
  szmf: [2.5 / 10, 7.5 / 10],
  hzmf: [2.5 / 10, 7.5 / 10],
  zsigmf: [-20, 5 / 10],
  ssigmf: [20, 5 / 10],
  zlinemf: [2.5 / 10, 7.5 / 10],
  slinemf: [2.5 / 10, 7.5 / 10],
  hsmf: [2.5 / 10, 7.5 / 10],
  ssmf: [2.5 / 10, 7.5 / 10],
  gbellmf: [2 / 10, 4, 5 / 10],
  gaussmf: [1 / 10, 5 / 10],
  // 3D
  conemf: [
    [0.5, 0.5],
    [0.4, 0.4]
  ],
  pyrammf: [
    [0.5, 0.5],
    [0.4, 0.4]
  ],
  trappyrammf: [
    [0.5, 0.5],
    [0.2, 1],
    [0.4, 0.4]],
  gsigmf: [
    [0.5, 0.5],
    [0, 10]
  ],
  gbell3dmf: [
    [0.5, 0.5],
    [0.2, 0.2],
    [2.5, 2.5]
  ],
  gauss3dmf: [
    [0.5, 0.5],
    [0.25, 0.25]
  ],
  hyperbolmf: [
    [0.5, 0.5],
    [0.25, 0.25]
  ],
  ellipsmf: [
    [0.5, 0.5],
    [0.4, 0.4]
  ]
}

export const paramsLetters = {
  conemf: ['x0', 'h'],
  pyrammf: ['x0', 'h'],
  trappyrammf: ['x0', 'h', 'a'],
  gsigmf: ['x0', 'a'],
  gbell3dmf: ['x0', 'a', 'b'],
  gauss3dmf: ['x0', 'a'],
  hyperbolmf: ['x0', 'a'],
  ellipsmf: ['x0', 'a']
}
