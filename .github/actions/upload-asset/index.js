const fs = require('fs')
const path = require('path')
const core = require('@actions/core')
const github = require('@actions/github')

console.log(fs.readdirSync(path.join(__dirname, "../../..")))

const fileName = "destroyTF.darwin.tar.gz"

async function run() {

    const octokit = new github.GitHub(process.env.GITHUB_TOKEN);

    // TODO: get values from env.
    const releaseTagResponse = await octokit.repos.getReleaseByTag({
        owner: "rayhaanbhikha",
        repo: "destroyTF",
        tag: "v0.0.0"
    })

    console.log(releaseTagResponse)

    // octokit.repos.uploadReleaseAsset({
    //     file: fs.readFileSync(path.join(__dirname, "..", "..", "..", fileName)),
    //     Headers: {
    //         'content-type': 'application/zip'
    //     },
    //     name: fileName,
    //     url: 
    // })


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