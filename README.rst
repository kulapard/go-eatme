EatMe
=====

Утилита для выполнения массовых операций с вложенными репозиториями.


Установка
---------

::

    pip install eatme


Запуск
------

::

    eatme COMMAND

Дополнительные ключи: - ``--help`` - вывод справочной информации по
доступным параметрам - ``--version`` - вывод версии приложения и даты
последнего обновления


Тестирование
------------

Тесты (nosetests) + проверка покрытия (coverage) + статическая проверка
кода на наличие грубых ошибок (flake8 + pylint) с помощю `tox`_:

::

    tox --skip-missing-interpreters --recreate

Тщательная статическая проверка кода:

::

    tox -e check --skip-missing-interpreters --recreate

.. _tox: https://pypi.python.org/pypi/tox

