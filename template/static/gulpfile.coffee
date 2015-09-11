gulp    = require 'gulp'
uglify = require 'gulp-uglify'
plumber = require 'gulp-plumber'


gulp.task 'default', ['build']
gulp.task 'build', [
  'build:bundle'
]

sass = require('gulp-sass')

gulp.task 'build:css', ->
  gulp
    .src('src/style/*.scss')
    .pipe plumber()
    .pipe(sass())
    .pipe(gulp.dest('dist/style/'))

# gulp.task 'compress:js', ['build:bundle'], ->
#   gulp.src 'temp/**/*.js'
#       .pipe uglify()
#       .pipe gulp.dest 'dist/'


through2 = require 'through2'
browserify = require 'browserify'
babelify = require 'babelify'
B_CONF = {
  debug: true
  basedir: './src/script/'
  shim: {
    lodash: {
      path: './node_modules/lodash/index.js'
      exports: 'lodash'
    }
    react: {
      path: './node_modules/react/dist/react-with-addons.js'
      exports: 'react'
    }
  }
}
makeBabel = () ->
  return through2.obj((file, enc, next) ->
    br = browserify(file.path, B_CONF)
      .transform 'coffeeify'
      .transform(babelify.configure({stage: 2}))
    br.bundle(
      (err, res) ->
        console.log file.path
        if err
          return next(err)
        file.contents = res
        next(null, file)
    )
  )

# via: https://github.com/substack/node-browserify/issues/1198
gulp.task 'build:js', ->
  gulp.src('./src/entry.coffee')
    .pipe(makeBabel())
    .pipe(gulp.dest('dist/js/'))


gulp.task 'default', ['build']


gulp.task 'watch', ['build'], ->
  gulp.watch 'src/**/*', ['build:bundle']
