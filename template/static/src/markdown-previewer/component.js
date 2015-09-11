md2react = require 'md2react'


###*
* @param {string} txt
* @return {Array}
###
_renderMd = (txt) ->
  try
    md2react txt,
      gfm: true
      breaks: true
      tables: true
  catch e
    console.warn 'mark down parse error', e
    []



module.exports =
  MarkdownPreviewerComponent = React.createClass
    mixins: [Arda.mixin]

    render: () ->
      console.log 'MarkdownPreviewerComponent render', @props
      template = require './view'
      template(_renderMd(@props.content), @props.addtionalClass)


