// 原型模式则不同，原型模式是先创建好一个原型对象，然后通过clone这个原型对象来创建新的对象，
// 这样就免去了重复的初始化工作，系统仅需内存拷贝即可

<?php
// interface和abstract，如何选择
abstract class Prototype {
	abstract public function copy();
	public function printInfo() {
		printf("%s\n", $this->info);
	}
}

class Student extends Prototype {
	public $info;
	public function __construct($info) {
		$this->info = $info;
	}

	public function setInfo($info) {
		$this->info = $info;
	}

	public function copy() {
		return clone $this;  // 浅拷贝
	}

	public function __clone() {
		$this->info = "HelloWorld3";
	}
}

$demo1 = new Student("HelloWorld1");
$demo1->printInfo();
$demo2 = $demo1->copy();
$demo1->printInfo();
$demo2->printInfo();

