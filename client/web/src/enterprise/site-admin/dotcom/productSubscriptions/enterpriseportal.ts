import type { PartialMessage } from '@bufbuild/protobuf'
import type { ConnectError, Transport } from '@connectrpc/connect'
import { defaultOptions, useMutation, useQuery } from '@connectrpc/connect-query'
import { createConnectTransport } from '@connectrpc/connect-web'
import { QueryClient, type UseMutationResult, type UseQueryResult } from '@tanstack/react-query'

import {
    getCodyGatewayAccess,
    getCodyGatewayUsage,
    updateCodyGatewayAccess,
} from './enterpriseportalgen/codyaccess-CodyAccessService_connectquery'
import type {
    GetCodyGatewayAccessResponse,
    GetCodyGatewayUsageResponse,
    UpdateCodyGatewayAccessRequest,
    UpdateCodyGatewayAccessResponse,
} from './enterpriseportalgen/codyaccess_pb'

/**
 * Use a shared QueryClient defined here and explicitly provided via
 * QueryClientProvider on only the components that need it, to avoid bleading
 * the QueryClientProvider to the site admin parent for now.
 *
 * Another problem is that @robert was unable to get QueryClientProvider working
 * even when placing it at various points the the tree.
 *
 * Overall this is only meant to be an interim integration, until Enterprise
 * Portal gets its own dedicated UI:
 * https://linear.app/sourcegraph/project/kr-p-enterprise-portal-user-interface-dadd5ff28bd8
 */
export const queryClient = new QueryClient({ defaultOptions })

/**
 * Use proxy that routes to a locally running Enterprise Portal at localhost:6081
 *
 * See cmd/frontend/internal/enterpriseportal/enterpriseportal_proxy.go
 */
const enterprisePortalLocal = createConnectTransport({
    baseUrl: '/.api/enterpriseportal/local',
})

/**
 * Use proxy that routes to the production Enterprise Portal at enterprise-portal.sourcegraph.com
 *
 * See cmd/frontend/internal/enterpriseportal/enterpriseportal_proxy.go
 */
const enterprisePortalDev = createConnectTransport({
    baseUrl: '/.api/enterpriseportal/dev',
})

/**
 * Use proxy that routes to the dev Enterprise Portal at enterprise-portal.sgdev.org
 *
 * See cmd/frontend/internal/enterpriseportal/enterpriseportal_proxy.go
 */
const enterprisePortalProd = createConnectTransport({
    baseUrl: '/.api/enterpriseportal/prod',
})

/**
 * Environment describes the Enteprise Portal instance to target. This can vary
 * per-subscription, as we currently use an unified management UI in Sourcegraph.com
 * as opposed to an Enterprise-Portal-specific UI, until:
 * https://linear.app/sourcegraph/project/kr-p-enterprise-portal-user-interface-dadd5ff28bd8
 *
 * 'local' is only valid in local dev.
 */
export type EnterprisePortalEnvironment =
    | 'prod' // enterprisePortalProd
    | 'dev' // enterprisePortalDev
    | 'local' // enterprisePortalLocal

const environments = new Map<EnterprisePortalEnvironment, Transport>([
    ['prod', enterprisePortalProd],
    ['dev', enterprisePortalDev],
    ['local', enterprisePortalLocal],
])

function mustGetEnvironment(env: EnterprisePortalEnvironment): Transport {
    if (env === 'local' && window.context.deployType !== 'dev') {
        throw new Error(`Environment ${env} not allowed outside local dev`)
    }
    const transport = environments.get(env)
    if (transport) {
        return transport
    }
    throw new Error(`Environment ${env} not configured`)
}

/**
 * Retrieves Cody Gateway usage for a given subscription.
 * @param env Enterprise Portal environment to target.
 * @param subscriptionUUID Enterprise Subscription UUID.
 */
export function useGetCodyGatewayUsage(
    env: EnterprisePortalEnvironment,
    subscriptionUUID: string
): UseQueryResult<GetCodyGatewayUsageResponse, ConnectError> {
    return useQuery(
        getCodyGatewayUsage,
        {
            query: { value: subscriptionUUID, case: 'subscriptionId' },
        },
        { transport: mustGetEnvironment(env) }
    )
}

export function useGetCodyGatewayAccess(
    env: EnterprisePortalEnvironment,
    subscriptionUUID: string
): UseQueryResult<GetCodyGatewayAccessResponse, ConnectError> {
    return useQuery(
        getCodyGatewayAccess,
        {
            query: { value: subscriptionUUID, case: 'subscriptionId' },
        },
        { transport: mustGetEnvironment(env) }
    )
}

export function useUpdateCodyGatewayAccess(
    env: EnterprisePortalEnvironment
): UseMutationResult<
    UpdateCodyGatewayAccessResponse,
    ConnectError,
    PartialMessage<UpdateCodyGatewayAccessRequest>,
    unknown
> {
    return useMutation(updateCodyGatewayAccess, {
        transport: mustGetEnvironment(env),
    })
}
