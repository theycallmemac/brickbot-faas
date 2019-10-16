# Literally taken from https://github.com/cpssd-students/steely/blob/master/steely/plugins/roomcheck.py 
# so Tom Doyle gets the credit for this one not me, I just modified it to work as a command line tool

import datetime
import sys
import requests
import bs4

COMMAND = 'room'
BASEURL = 'https://www101.dcu.ie/timetables/feed.php'
def parse_room_number(room_number):
    room_number = room_number.upper()
    campus_codes = ('GLA.', 'SPD.')
    default_campus = campus_codes[0]
    if not any(campus in room_number for campus in campus_codes):
        room_number = default_campus + room_number
    return room_number


def get_booking(room, baseurl=None):
    baseurl = baseurl or BASEURL
    params = build_request_parameters(room)
    response = requests.get(baseurl, params=params)
    current_soup = bs4.BeautifulSoup(response.text, "lxml")
    elements = current_soup.select('tr')
    return elements[14].getText().strip(), response.url


def build_request_parameters(room_number):
    now = datetime.datetime.now()
    return {'room': room_number,
            'week': academic_week_number(now),
            'hour': academic_hour(now),
            'day':  now.isoweekday(),
            'template': 'location'}


def academic_week_number(date):
    iso_year, iso_week_number, iso_weekday = date.isocalendar()
    academic_start_week = 36
    if iso_week_number >= academic_start_week:
        academic_offset = -academic_start_week
    else:
        academic_offset = 52 - academic_start_week
    academic_week = iso_week_number + academic_offset
    return academic_week


def academic_hour(date):
    hour = date.hour
    if hour < 8 or hour > 21:
        raise ValueError('invalid academic hour')
    minute = date.minute
    if minute > 30:
        hour += 1
    return (hour - 8) + (hour - 7)

def main():
    try:
        room_number = parse_room_number(sys.argv[1].upper())
    except IndexError:
        print('no room was supplied')
        sys.exit(0)
    try:
        booking, link = get_booking(room_number)
    except ValueError:
        print('building is closed at this time')
        sys.exit(0)
    if booking:
        print(f'{room_number} is not free, there is\n\n' +
                     f'{booking}')
    else:
        print(f'{room_number} is currently free')
    print(link)

if __name__ == "__main__":
    main()
