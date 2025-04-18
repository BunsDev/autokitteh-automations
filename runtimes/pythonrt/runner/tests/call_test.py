import re
import sys
from pathlib import Path
from time import sleep
from types import ModuleType
from unittest.mock import MagicMock

import pytest
from autokitteh import activity, AttrDict
from conftest import clear_module_cache, workflows

import call
import loader


def test_sleep():
    runner = MagicMock()
    ak_call = call.AKCall(runner, Path("/tmp"))
    mod_name = "program"
    clear_module_cache(mod_name)

    mod = loader.load_code(workflows.sleeper, ak_call, mod_name)
    ak_call.set_module(mod)
    event = {"type": "login", "user": "puss"}
    mod.handler(event)
    assert runner.syscalls.ak_sleep.call_count == 2


def test_sleep_activity():
    comm = MagicMock()
    ak_call = call.AKCall(comm, Path("/tmp"))
    ak_call.in_activity = True
    ak_call(sleep, 0.1)

    assert comm.send_call.call_count == 0


def test_call_non_func():
    comm = MagicMock()
    ak_call = call.AKCall(comm, Path("/tmp"))
    with pytest.raises(ValueError):
        ak_call("hello")


def test_should_run_as_activity(monkeypatch):
    code_dir = Path(__file__).parent
    mod_name = "ak_test_module_name"
    mod = ModuleType(mod_name)
    mod.__file__ = str(code_dir / (mod_name + ".py"))

    monkeypatch.setitem(sys.modules, mod_name, mod)

    ak_call = call.AKCall(None, code_dir)

    def fn():
        pass

    fn.__module__ = "/tmp"  # Not from user dir

    # loading
    assert not ak_call.should_run_as_activity(fn)

    # after loading
    ak_call.set_module(mod)
    assert ak_call.should_run_as_activity(fn)

    # Function from same module
    fn.__module__ = mod_name
    assert not ak_call.should_run_as_activity(fn)

    # Function from user code dir
    fn.__file__ = str(code_dir / "other.py")
    assert not ak_call.should_run_as_activity(fn)

    # Marked activity
    act_fn = activity(fn)
    assert ak_call.should_run_as_activity(act_fn)

    # In activity
    ak_call.in_activity = True
    assert not ak_call.should_run_as_activity(act_fn)
    ak_call.in_activity = False

    # Deterministic
    assert not ak_call.should_run_as_activity(re.compile)

    # AutoKitteh function
    evt = AttrDict()
    assert not ak_call.should_run_as_activity(evt.get)


def test_is_module_func(monkeypatch: pytest.MonkeyPatch):
    mod_name = "handler"
    clear_module_cache(mod_name)
    code_dir = workflows.multi_file
    monkeypatch.syspath_prepend(str(code_dir))

    runner = MagicMock()
    ak_call = call.AKCall(runner, code_dir)
    mod = loader.load_code(code_dir, ak_call, mod_name)
    ak_call.set_module(mod)

    assert ak_call.is_module_func(mod.on_event)  # Same handler file
    assert ak_call.is_module_func(mod.hlog.info)  # Same directory
    assert not ak_call.is_module_func(mod.json.dump)


def named_fn():
    pass


class Inc:
    def __init__(self, n):
        self._n = n

    def __call__(self, v):
        return v + self._n

    def n(self):
        return self._n


name_cases = [
    (named_fn, "call_test.named_fn"),
    (Inc(3), "call_test.Inc"),
    (Inc(3).n, "call_test.Inc.n"),
]


@pytest.mark.parametrize("fn, name", name_cases)
def test_full_func_name(fn, name):
    assert call.full_func_name(fn) == name
