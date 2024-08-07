import { Suspense, type FC } from 'react'

import { mdiClose } from '@mdi/js'

import { CodyLogo } from '@sourcegraph/cody-ui'
import { lazyComponent } from '@sourcegraph/shared/src/util/lazyComponent'
import { Alert, Button, H4, Icon, LoadingSpinner, ProductStatusBadge } from '@sourcegraph/wildcard'

import styles from './NewCodySidebar.module.scss'

const LazyCodySidebarWebChat = lazyComponent(() => import('./NewCodySidebarWebChat'), 'NewCodySidebarWebChat')

export interface Repository {
    id: string
    name: string
}

interface NewCodySidebarProps {
    filePath: string | undefined
    repository: Repository
    isAuthorized: boolean
    onClose: () => void
}

export const NewCodySidebar: FC<NewCodySidebarProps> = props => {
    const { repository, filePath, isAuthorized, onClose } = props

    return (
        <div className={styles.root}>
            <div className={styles.header}>
                <div />
                <span className={styles.headerLogo}>
                    <CodyLogo />
                    Cody
                    <div className="ml-2">
                        <ProductStatusBadge status="beta" />
                    </div>
                </span>

                <Button variant="icon" aria-label="Close" onClick={onClose}>
                    <Icon aria-hidden={true} svgPath={mdiClose} />
                </Button>
            </div>

            {isAuthorized && (
                <Suspense
                    fallback={
                        <div className="flex flex-1 align-items-center m-2">
                            <LoadingSpinner className="mr-2" /> Loading Cody client
                        </div>
                    }
                >
                    <LazyCodySidebarWebChat filePath={filePath} repository={repository} />
                </Suspense>
            )}

            {!isAuthorized && (
                <Alert variant="info" className="m-3">
                    <H4>Cody is only available to signed-in users</H4>
                    Sign in to get access to use Cody
                </Alert>
            )}
        </div>
    )
}
