const fs = require('fs')
const path = require('path')
const core = require('@actions/core')
const github = require('@actions/github')

const { GITHUB_REPOSITORY, GITHUB_TOKEN, GITHUB_REF } = process.env

function getFileSizeInBytes(pathToFile) {
    return fs.statSync(pathToFile)["size"]
}

async function run() {

    try {
        const [owner, repo] = GITHUB_REPOSITORY.split("/")
        const githubRef = GITHUB_REF.split("/")
        const tag = githubRef[githubRef.length - 1]

        const fileName = `destroyTF.darwin-${tag}.tar.gz`
        const pathToFile = path.join(__dirname, "..", "..", "..", fileName)

        const octokit = new github.GitHub(GITHUB_TOKEN);

        // TODO: get values from env.
        const { data } = await octokit.repos.getReleaseByTag({
            owner,
            repo,
            tag
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