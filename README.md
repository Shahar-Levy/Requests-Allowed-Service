Given an IP address and a white list of countries, this service aims to return an indicator if the IP address is within the listed countries. I grabbed the countries and IP ranges from csv files at https://dev.maxmind.com/geoip/geoip2/geolite2/

This service is incomplete. It still needs to:
1) Take input (ip, whitelist) from the request header
2) Check against IPs instead of an IP range. The CSV files from maxmind give an IP range and the logic to find if an IP is in an IP range is incomplete