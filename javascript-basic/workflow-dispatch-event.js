// Octokit.js
// https://github.com/octokit/core.js#readme
const octokit = new Octokit({
  auth: 'YOUR-TOKEN'
})

await octokit.request('POST /repos/tainguyenbp/learn-programming/actions/workflows/gha-security-checks.yml/dispatches', {
  owner: 'OWNER',
  repo: 'REPO',
  workflow_id: 'WORKFLOW_ID',
  ref: 'topic-branch',
  inputs: {
    name: 'Mona the Octocat',
    home: 'San Francisco, CA'
  },
  headers: {
    'X-GitHub-Api-Version': '2022-11-28'
  }
})


detection-forensics-cheatsheet
by defensive-security
