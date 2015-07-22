kd                   = require 'kd'
globals              = require 'globals'
KDTabPaneView        = kd.TabPaneView
KDCustomHTMLView     = kd.CustomHTMLView
AdminMainTabPaneView = require './views/adminmaintabpaneview'


module.exports = class AdminAppView extends kd.ModalView

  constructor: (options = {}, data) ->

    options.testPath = 'groups-admin'
    data           or= kd.singletons.groupsController.getCurrentGroup()

    super options, data

    @addSubView @nav  = new kd.TabHandleContainer
      cssClass: 'AppModal-nav'

    @addSubView @tabs = new AdminMainTabPaneView
      tabHandleContainer: @nav
    , data

    @nav.unsetClass 'kdtabhandlecontainer'

    @setListeners()


  setListeners: ->

    @on 'groupSettingsUpdated', (group)->
      @setData group
      @createTabs()

    @tabs.on 'PaneAdded', (pane) ->
      { tabHandle } = pane
      tabHandle.setClass 'AppModal-navItem'
      tabHandle.unsetClass 'kdtabhandle'


  viewAppended: ->

    super

    group = kd.getSingleton('groupsController').getCurrentGroup()
    group?.canEditGroup (err, success) =>
      if err or not success
        {entryPoint} = globals.config
        kd.singletons.router.handleRoute '/Activity', { entryPoint }
      else
        @createTabs()


  createTabs: ->

    data         = @getData()
    {tabData}    = @getOptions()
    currentGroup = kd.singletons.groupsController.getCurrentGroup()

    items = []

    for own sectionKey, section of tabData

      if sectionKey is 'koding' and currentGroup.slug isnt 'koding'
        continue

      for item in section.items
        items.push item

        if item.subTabs
          for subTab in item.subTabs
            subTab.parentTabTitle = item.title
            items.push subTab


    @tabs.on 'PaneDidShow', (pane) ->
      return  if pane._isViewAdded

      slug       = pane.getOption 'slug'
      action     = pane.getOption 'action'
      identifier = pane.getOption 'identifier'
      targetItem = viewClass: KDCustomHTMLView

      for item in items
        if item.action is action
          targetItem = item
          break
        else if item.slug is slug
          targetItem = item
          break

      { viewClass } = targetItem

      pane._isViewAdded = yes
      pane.setMainView new viewClass
        cssClass : slug or action
        delegate : this
        action   : action
      , data

    items.forEach (item, i) =>

      { slug, title, action } = item
      name           = slug or action
      hiddenHandle   = if action then yes
      parentTabTitle = item.parentTabTitle or null

      pane = new KDTabPaneView { name, slug, action, hiddenHandle, title, parentTabTitle }
      @tabs.addPane pane, i is 0

    @emit 'ready'


  search: (searchValue) ->

    if @tabs.getActivePane().name is 'Invitations'
      pane = @tabs.getActivePane()
    else
      pane = @tabs.getPaneByName 'Members'
      @tabs.showPane pane

    pane?.mainView?.emit 'SearchInputChanged', searchValue
