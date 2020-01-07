const fs = require('fs')
const path = require('path')
const core = require('@actions/core')
const github = require('@actions/github')

console.log(fs.readdirSync(path.join(__dirname, "../../..")))

const fileName = "destroyTF.darwin.tar.gz"

async function run() {

    const octokit = new github.GitHub(process.env.GITHUB_TOKEN);

    // TODO: get values from env.
    const { data } = await octokit.repos.getReleaseByTag({
        owner: "rayhaanbhikha",
        repo: "destroyTF",
        tag: "v0.0.0"
    })

    const response = await octokit.repos.uploadReleaseAsset({
        file: fs.readFileSync(path.join(__dirname, "..", "..", "..", fileName)),
        Headers: {
            'content-type': 'application/zip'
        },
        name: fileName,
        url: data.upload_url
    })

    console.log(response)
}

run();