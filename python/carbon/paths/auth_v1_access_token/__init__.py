# do not import all endpoints into this module because that uses a lot of memory and stack frames
# if you need the ability to import all endpoints from this module, import them with
# from carbon.paths.auth_v1_access_token import Api

from carbon.paths import PathValues

path = PathValues.AUTH_V1_ACCESS_TOKEN