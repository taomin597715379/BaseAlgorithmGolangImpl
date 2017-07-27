// 这里使用php去配合编写设计模式Demo代码
// Factory这种模式提供了一个让子类实例化哪一个类，工厂模式使得一个类的实例化可以延迟到其子类
// 工厂和产品是什么关系，一对多的关系吗？还是一对一的关系？
// 工厂和产品是一一对应的，但产品种类可以任意，因为只需在产品类里面增加产品的生产方式

<?php

interface Creater {
	public function factoryMethod();
}

class aFactory implements Creater {
	public function factoryMethod() {
		return new aProduct();
	}
}

class bFactory implements Creater {
	public function factoryMethod() {
		return new bProduct();
	}
}

class aProduct {
	public function productMySelf1() {
		echo "aProduct1" . PHP_EOL;
	}

	public function productMySelf2() {
		echo "aProduct2" . PHP_EOL;
	}
}

class bProduct {
	public function productMySelf1() {
		echo "bProduct1" . PHP_EOL;
	}

	public function productMySelf2() {
		echo "bProduct2" . PHP_EOL;
	}
}

$creater = new aFactory();
$creater->factoryMethod()->productMySelf1();
$creater = new bFactory();
$creater->factoryMethod()->productMySelf2();



