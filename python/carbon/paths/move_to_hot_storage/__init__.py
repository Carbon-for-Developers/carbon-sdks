# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from carbon.paths.move_to_hot_storage import Api

from carbon.paths import PathValues

path = PathValues.MOVE_TO_HOT_STORAGE