class ControllerError(Exception):
    """Base class for all controller errors"""


class CommandError(ControllerError):
    """Base class for all command errors"""


class ShellCommandError(CommandError):
    def __init__(self, cmd: str, err_msg: str, data: bytes):
        self.cmd = cmd
        self.err_msg = err_msg
        super().__init__(
            f'Failed to execute shell command "{cmd}": {err_msg} ; {data}')


class DownloadFileError(CommandError):
    def __init__(self, remote_path: str, err_msg: str):
        self.remote_path = remote_path
        self.err_msg = err_msg
        super().__init__(f'Failed to download file "{remote_path}": {err_msg}')


class UploadFileError(CommandError):
    def __init__(self, remote_path: str, err_msg: str):
        self.remote_path = remote_path
        self.err_msg = err_msg
        super().__init__(f'Failed to upload file "{remote_path}": {err_msg}')
