// 这里使用php去配合编写设计模式Demo代码
// Singleton 是一种提供全局唯一共享的对象

<?php
class Singleton {
	private $name;
	private static $_instance = null;

	public function __construct() {
		
	}

	public static function getInstance() {
		if(is_null(self::$_instance)) {
			self::$_instance = new Singleton();
		}
		return self::$_instance;
	}

	public function __clone() {

	}

	public function setName($name) {
		$this->name = $name;
	}

	public function getName() {
		return $this->name;
	}
}

$client = Singleton::getInstance();
$client->setName("Hello");
var_dump($client->getName());