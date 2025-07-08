#!/bin/bash -x

# dmidecode is not available on ppc64le/s390x
if [[ "$(arch)" == "ppc64le" ]] || [[ "$(arch)" == "s390x" ]]; then
    echo "Non OCI architecture... exiting early"
    exit 0
fi

chassis_asset_tag="$(dmidecode --string chassis-asset-tag)"
if [[ "${chassis_asset_tag}" != "OracleCloud.com" ]]; then
    echo "Not running in Oracle Cloud Infrastructure. Skipping."
    exit 0
fi

current_hostname="$(hostname)"
echo "Current hostname: ${current_hostname}"
if [[ "${current_hostname}" != "localhost.localdomain" ]]; then
    echo "Hostname has already been set. Skipping."
    exit 0
fi

OCI_HOSTNAME=/etc/hostname-oci
until [[ -s $OCI_HOSTNAME ]]; do
    /usr/bin/curl -s -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/hostname -o $OCI_HOSTNAME
done

echo "Setting hostname to $(cat $OCI_HOSTNAME)"

cat $OCI_HOSTNAME > /proc/sys/kernel/hostname
