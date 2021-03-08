const colors = require('windicss/colors')
const typography = require('windicss/plugin/typography')

module.exports = {
  darkMode: 'class',
  plugins: [typography],
  theme: {
    extend: {
      colors: {
        teal: colors.teal
      }
    },
    listStyleType: {
      none: 'none',
      disc: 'disc',
      decimal: 'decimal',
      lowerlatin: 'lower-latin'
    }
  }
}
