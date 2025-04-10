-- Common class packages
INSERT INTO packages (package_name, price, startTime, endTime)
VALUES
    ('1 Hour', 400, NULL, NULL),
    ('2 Hour', 600, NULL, NULL),
    ('3 Hour', 800, NULL, NULL),
    ('5 Hour', 1200, NULL, NULL),
    ('Day Package', 2000, '08:00', '20:00'),
    ('Night Package', 1500, '22:00', '06:00'),
    ('Morning Package', 1500, '06:00', '12:00');
