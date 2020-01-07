const fs = require('fs')
const path = require('path')
const core = require('@actions/core')
const github = require('@actions/github')
console.log("hello running a javascript action")


console.log(process.env.GITHUB_TOKEN)
console.log(fs.readdirSync(path.join(__dirname, "../../..")))

async function run() {
    // This should be a token with access to your repository scoped in as a secret.
    // The YML workflow will need to set myToken with the GitHub Secret Token
    // myToken: ${{ secrets.GITHUB_TOKEN }}
    // https://help.github.com/en/actions/automating-your-workflow-with-github-actions/authenticating-with-the-github_token#about-the-github_token-secret
    const githubToken = core.getInput('GITHUB_TOKEN');
    console.log(githubToken)
    // const octokit = new github.GitHub(githubToken);

    // const { data: pullRequest } = await octokit.pulls.get({
    //     owner: 'octokit',
    //     repo: 'rest.js',
    //     pull_number: 123,
    //     mediaType: {
    //       format: 'diff'
    //     }
    // });

    // console.log(pullRequest);
}

run();