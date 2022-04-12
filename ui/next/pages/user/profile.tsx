// @ts-nocheck
import React, { useEffect, useState } from 'react'
import get from 'lodash/get'
import { Layout } from 'components'
import Welcome from 'components/Profile/Welcome'
import Profile from 'components/Profile/Profile'
import Personal from 'components/Profile/Personal'
import Notifications from 'components/Profile/Notifications'
import withAuthSync from 'components/Private'
import ory from '../../pkg/sdk'
import { AxiosError } from 'axios'

function ProfileContent() {
  const [session, setSession] = useState<string>(
    'No valid Ory Session was found.\nPlease sign in to receive one.',
  )
  const [hasSession, setHasSession] = useState<boolean>(false) // eslint-disable-line

  useEffect(() => {
    ory
      .toSession()
      .then(({ data }) => {
        setSession(JSON.stringify(data, null, 2))
        setHasSession(true)
      })
      .catch((err: AxiosError) =>
        // Something else happened!
        Promise.reject(err),
      )
  })

  return (
    <Layout>
      <Welcome
        nickname={get(session, 'kratos.identity.traits.name.first', 'default')}
      />

      <Profile />

      <div className="hidden sm:block" aria-hidden="true">
        <div className="py-5">
          <div className="border-t border-gray-200" />
        </div>
      </div>

      <Personal session={session} />

      <div className="hidden sm:block" aria-hidden="true">
        <div className="py-5">
          <div className="border-t border-gray-200" />
        </div>
      </div>

      <Notifications />
    </Layout>
  )
}

export default withAuthSync(ProfileContent)
