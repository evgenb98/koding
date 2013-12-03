class WebTermController extends AppController

  KD.registerAppClass this,
    name         : "Terminal"
    title        : "Terminal"
    navItem      :
      title      : "Terminal"
      order      : 41
      path       : "/Terminal"
    route        :
      slug       : "/:name?/Terminal"
      handler    : ({params:{name}, query})->
        # KD.utils.wait 800, ->
        router = KD.getSingleton 'router'
        router.openSection "Terminal", name, query
    multiple     : yes
    hiddenHandle : no
    menu         :
      width      : 250
      items      : [
        {title: "customViewAdvancedSettings"}
      ]
    behavior     : "application"
    preCondition :
      condition  : (options, cb)->

        {vmName} = options
        vmController = KD.getSingleton 'vmController'
        vmController.fetchDefaultVmName (defaultVmName)->
          vmName or= defaultVmName
          return cb no  unless vmName
          vmController.info vmName, KD.utils.getTimedOutCallback (err, vm, info)->
            cb  info?.state is 'RUNNING', {vmName, info}
          , ->
            cb  no
            unless KD.isGuest()
              KD.logToExternal "failed to fetch vminfo, couldn't open terminal"
          , 2500

      failure     : (options, cb)->
        KD.getSingleton("vmController").askToTurnOn
          appName : 'Terminal'
          vmName  : options.vmName
          state   : options.info.state
        , cb

  constructor:(options = {}, data)->
    params              = options.params or {}
    {joinUser, session} = params
    vmName              = params.vmName  or KD.getSingleton("vmController").defaultVmName
    options.view        = new WebTermAppView { vmName, joinUser, session }
    options.appInfo     =
      title             : "Terminal on #{vmName}"
      cssClass          : "webterm"

    super options, data

  handleQuery: (query) ->
    @getView().ready =>
      @getView().handleQuery query


WebTerm = {}
