const fs = require('fs')
const path = require('path')
const core = require('@actions/core')
const github = require('@actions/github')

console.log(process.env)

const fileName = "destroyTF.darwin.tar.gz"
const pathToFile = path.join(__dirname, "..", "..", "..", fileName)

function getFileSizeInBytes(pathToFile) {
    return fs.statSync(pathToFile)["size"]
}

async function run() {

    try {
        const octokit = new github.GitHub(process.env.GITHUB_TOKEN);

        // TODO: get values from env.
        const { data } = await octokit.repos.getReleaseByTag({
            owner: "rayhaanbhikha",
            repo: "destroyTF",
            tag: "v0.0.0"
        })

        await octokit.repos.uploadReleaseAsset({
            file: fs.readFileSync(pathToFile),
            headers: {
                'content-type': 'application/zip',
                'content-length': getFileSizeInBytes(pathToFile)
            },
            name: fileName,
            url: data.upload_url
        })
    } catch (err) {
        core.error(err)
        core.setFailed(err.message)
    }
}

run();