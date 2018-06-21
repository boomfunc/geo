import os
import sys

import pygeoip


def get_geo_record(ip):
	try:
		# get geo database and
		# get record by ip address ( dict )
		geo_record = pygeoip.GeoIP(
			os.path.join(
				os.getcwd(),
				'geoip_city.dat'
			)
		).record_by_addr(ip)
	except (pygeoip.GeoIPError, IOError, AttributeError):
		geo_record = None

	return geo_record or {}


if __name__ == '__main__':
	ip = sys.stdin.read()
	sys.stdout.write('NDJKNDJKDNJKDNDJKNDJKNDJKDNJKDNDJKNDJKDNJKDNDNJK', ip)
