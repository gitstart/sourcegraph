import '../platform/polyfills'

import { ShortcutProvider } from '@slimsag/react-shortcuts'
import { VSCodeProgressRing } from '@vscode/webview-ui-toolkit/react'
import * as Comlink from 'comlink'
import React, { useMemo } from 'react'
import { render } from 'react-dom'

import { wrapRemoteObservable } from '@sourcegraph/shared/src/api/client/api/common'
import { AnchorLink, setLinkComponent, useObservable, WildcardThemeContext, Tooltip } from '@sourcegraph/wildcard'

import { ExtensionCoreAPI } from '../../contract'
import { createEndpointsForWebToNode } from '../comlink/webviewEndpoint'
import { createPlatformContext, WebviewPageProps } from '../platform/context'
import { adaptMonacoThemeToEditorTheme } from '../theming/monacoTheme'
import { adaptSourcegraphThemeToEditorTheme } from '../theming/sourcegraphTheme'

import { SearchHomeView } from './SearchHomeView'
import { SearchResultsView } from './SearchResultsView'

import './index.module.scss'
import { searchPanelAPI } from './api'

const vsCodeApi = window.acquireVsCodeApi()

const { proxy, expose } = createEndpointsForWebToNode(vsCodeApi)

Comlink.expose(searchPanelAPI, expose)

export const extensionCoreAPI: Comlink.Remote<ExtensionCoreAPI> = Comlink.wrap(proxy)

const themes = adaptSourcegraphThemeToEditorTheme()
adaptMonacoThemeToEditorTheme()

extensionCoreAPI.panelInitialized(document.documentElement.dataset.panelId!).catch(() => {
    // noop (TODO?)
})

const platformContext = createPlatformContext(extensionCoreAPI)

setLinkComponent(AnchorLink)

const Main: React.FC = () => {
    const state = useObservable(useMemo(() => wrapRemoteObservable(extensionCoreAPI.observeState()), []))

    const authenticatedUser = useObservable(
        useMemo(() => wrapRemoteObservable(extensionCoreAPI.getAuthenticatedUser()), [])
    )

    const instanceURL = useObservable(useMemo(() => wrapRemoteObservable(extensionCoreAPI.getInstanceURL()), []))

    const theme = useObservable(themes)

    const settingsCascade = useObservable(
        useMemo(() => wrapRemoteObservable(extensionCoreAPI.observeSourcegraphSettings()), [])
    )
    // Do not block rendering on settings unless we observe UI jitter

    // TODO: If init is taking too long, show a message.
    // Also check if anything has errored out.

    // If any of the remote values have yet to load.
    const initialized =
        state !== undefined &&
        authenticatedUser !== undefined &&
        instanceURL !== undefined &&
        theme !== undefined &&
        settingsCascade !== undefined

    if (!initialized) {
        return <VSCodeProgressRing />
    }

    const webviewPageProps: WebviewPageProps = {
        extensionCoreAPI,
        platformContext,
        theme,
        authenticatedUser,
        settingsCascade,
        instanceURL,
    }

    if (state?.status === 'context-invalidated') {
        // TODO context-invalidated state
        return null
    }

    // Render SearchHomeView until the user submits a search.
    if (state.context.submittedSearchQueryState === null) {
        return <SearchHomeView {...webviewPageProps} context={state.context} />
    }

    return (
        <SearchResultsView
            {...webviewPageProps}
            context={{
                ...state.context,
                submittedSearchQueryState: state.context.submittedSearchQueryState,
            }}
        />
    )
}

render(
    <ShortcutProvider>
        <WildcardThemeContext.Provider value={{ isBranded: true }}>
            <Main />
            <Tooltip key={1} className="sourcegraph-tooltip" />
        </WildcardThemeContext.Provider>
    </ShortcutProvider>,
    document.querySelector('#root')
)
