class KDSplitViewPanel extends KDScrollView
  constructor:(options,data)->
    # options = $.extend
    #   ownScrollBars : yes
    # ,options
    super options,data
    @isVertical = @options.type.toLowerCase() is "vertical"
    {@size,@minimum,@maximum,@index} = @options

  _getSize:->if @isVertical then @getWidth() else @getHeight()

  _setSize:(size)->
    if @_wouldResize size
      size = 0 if size < 0
      if @isVertical then @setWidth size else @setHeight size
      @parent.sizes[@index] = @size = size
      @parent.emit "PanelDidResize", panel: @
      @emit "PanelDidResize", newSize : size
      size
    else
      no

  _wouldResize:(size)->
    @minimum ?= -1
    @maximum ?= 99999
    # log size,@minimum,@maximum if @parent.options.domId is "content-area-split-view"
    if size > @minimum and size < @maximum
      # log size,@parent.options.domId
      yes
    else
      if size < @minimum
        @parent._panelReachedMinimum @index
      else if size > @maximum
        @parent._panelReachedMaximum @index
      no

  _setOffset:(offset)->
    offset = 0 if offset < 0
    if @isVertical then @$().css(left : offset) else @$().css(top : offset)
    @parent.panelsBounds[@index] = offset

  _getOffset:->
    if @isVertical then @getRelativeX() else @getRelativeY()

  _animateTo:(size,offset,callback)=>
    if "undefined" is typeof callback and "function" is typeof offset then callback = offset
    callback or= noop

    panel = @
    d     = panel.parent.options.duration
    cb    = ()->
      # setTimeout do ->
      newSize = panel._getSize()
      panel.parent.sizes[panel.index] = panel.size = newSize
      panel.parent.emit "PanelDidResize", panel: panel
      panel.emit "PanelDidResize", newSize : newSize
      callback.call panel
      # ,100


    properties = {}
    size = 0 if size < 0
    if panel.isVertical
      properties.width  = size
      properties.left   = offset if offset?
    else
      properties.height = size
      properties.top    = offset if offset?

    options =
      duration : d
      complete : cb
      # step     : (newSize)-> panel.parent.handleEvent {
      #   type : "PanelIsBeingResized"
      #   panel
      #   newSize
      # }

    panel.$().stop()
    panel.$().animate properties,options
