<?php declare(strict_types=1);

use Phroute\Phroute\RouteCollector;
use Phroute\Phroute\Dispatcher;

require __DIR__ . '/vendor/autoload.php';

$router = new RouteCollector();

$router->get('/hello/{name}', function(string $name){
    return 'Hello ' . $name;
});

$dispatcher = new Dispatcher($router->getData());

$response = $dispatcher->dispatch($_SERVER['REQUEST_METHOD'], parse_url($_SERVER['REQUEST_URI'], PHP_URL_PATH));
print $response;