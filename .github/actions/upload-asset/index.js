const fs = require('fs')
const path = require('path')
const core = require('@actions/core')
const github = require('@actions/github')

console.log(fs.readdirSync(path.join(__dirname, "../../..")))

const fileName = "destroyTF.darwin.tar.gz"
const pathToFile = path.join(__dirname, "..", "..", "..", fileName)

function getFileSizeInBytes(pathToFile) {
    return fs.statSync(pathToFile)["size"]
}

async function run() {

    const octokit = new github.GitHub(process.env.GITHUB_TOKEN);

    // TODO: get values from env.
    const { data } = await octokit.repos.getReleaseByTag({
        owner: "rayhaanbhikha",
        repo: "destroyTF",
        tag: "v0.0.0"
    })

    const response = await octokit.repos.uploadReleaseAsset({
        file: fs.readFileSync(pathToFile),
        headers: {
            'content-type': 'application/zip',
            'content-length': getFileSizeInBytes(pathToFile)
        },
        name: fileName,
        url: data.upload_url
    })

    console.log(response)
}

run();