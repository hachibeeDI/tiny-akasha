gulp    = require 'gulp'
shell   = require 'gulp-shell'
coffee  = require 'gulp-coffee'
uglify = require('gulp-uglify')
# sass    = require 'gulp-sass'
plumber = require 'gulp-plumber'
# changed  = require 'gulp-changed'

gulp.task 'default', ['build']
gulp.task 'build', [
  'build:bundle'
]

gulp.task 'build:coffee', ->
  gulp.src('src/**/*.coffee')
    .pipe(coffee())
    .pipe(gulp.dest('temp'))



sass = require('gulp-sass')

gulp.task 'build:css', ->
  gulp
    .src('src/style/*.scss')
    .pipe plumber()
    .pipe(sass())
    .pipe(gulp.dest('dist/style/'))

gulp.task 'compress:js', ['build:coffee', 'build:bundle'], ->
  gulp.src 'temp/**/*.js'
      .pipe uglify()
      .pipe gulp.dest 'dist/'


# source = require 'vinyl-source-stream'
webpack = require('gulp-webpack')

gulp.task 'build:bundle', ['build:coffee', 'build:css'], shell.task [
  'webpack'
]


gulp.task 'default', ['build']


gulp.task 'watch', ['build'], ->
  gulp.watch 'src/**/*', ['build:bundle']
