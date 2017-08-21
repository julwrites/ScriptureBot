def get_name(name, last, alias):
    name_string = name
    if last:
        name_string += ' ' + last
    if alias:
        name_string += ' @' + alias
    return name_string
