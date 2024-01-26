import { logCommand } from './helpers'

const NO_LOG = { log: false }

export default function login(username, password): void {
  // WARNING: Whenever you change the code of the session setup, Cypress will throw an error:
  //   This session already exists. You may not create a new session with a previously used identifier.
  //   If you want to create a new session with a different setup function, please call cy.session() with
  //   a unique identifier other than...
  //
  // Temporarily uncomment the following line to clear the sessions if this happens
  // Cypress.session.clearAllSavedSessions()

  logCommand('login', { username }, username)

  cy.session(
    username,
    () => {
      cy.visit('/', NO_LOG)

      cy.contains('.btn', 'Log in', NO_LOG).click(NO_LOG)

      cy.origin(
        Cypress.env('OIDC_ORIGIN'),
        { args: { username, password, NO_LOG } },
        ({ username, password, NO_LOG }) => {
          cy.get('input[name="username"]', NO_LOG).type(username, NO_LOG)

          if (password) {
            cy.get('input[name="password"]', NO_LOG).type(password, NO_LOG)
          }

          cy.get(':submit', NO_LOG).click(NO_LOG)
        }
      )
    },
    {
      cacheAcrossSpecs: true,
    }
  )
}

declare global {
  namespace Cypress {
    interface Chainable {
      login(username: string, password: string): Chainable<void>
    }
  }
}
