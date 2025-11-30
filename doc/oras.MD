# How to deal with oras repository

# oras login with google sa token
cat registry-430716-d17542a2a5da.json | oras login oci://europe-west3-docker.pkg.dev/registry-430716/pottmeier --username _json_key --password-stdin

# push of a complete subdir
oras push europe-west3-docker.pkg.dev/registry-430716/pottmeier/sopstest:20251130 ./sopstest/

# checkout in current folder
oras pull .  europe-west3-docker.pkg.dev/registry-430716/pottmeier/sopstest:20251130


