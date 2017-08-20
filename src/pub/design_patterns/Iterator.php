// Iterator迭代器模型用来遍历标准模板库容器中的部分或全部元素
// 迭代器又称之为游标模型，提供访问容器的方法而不对外暴露细节

<?php
    interface IAggregate {
        public function createIterator();
    }

    class ConcreteAggregate implements IAggregate {
        public $workers;

        public function addElement($element) {
            $this->workers[] = $element;
        }

        public function getAt($index) {
            return $this->workers[$index];
        }

        public function current($index) {
        	return $this->workers[$index];
        }

        public function getLength() {
            return count($this->workers);
        }

        public function createIterator() {
            return new ConcreteIterator($this);
        }
    }

    interface IIterator {
        public function hasNext();
        public function next();
        public function current();
        public function reset();
    }

    class ConcreteIterator implements IIterator {
        public $collection;
        public $index;

        public function __construct($collection) {
            $this->collection = $collection;
            $this->index = 0;
        }

        public function current() {
        	return $this->collection->current($this->index);
        }

        public function reset() {
        	$this->index = 0;
        }

        public function hasNext() {
            if($this->index < $this->collection->getLength()) {
                return true;
            }
            else {
                return false;
            }
        }

        public function next() {
            $element = $this->collection->getAt($this->index);
            $this->index++;
            return $element;
        }
    }

    $farmerAggregate = new ConcreteAggregate();

    $farmerAggregate->addElement('SVC1');
    $farmerAggregate->addElement('SVC2');
    $farmerAggregate->addElement('SVC3');

    $iterator = $farmerAggregate->createIterator();

    while ($iterator->hasNext()) {
    	echo $iterator->current() . PHP_EOL;
        $element = $iterator->next();
        echo $element . PHP_EOL;
    }
    $iterator->reset();
    while ($iterator->hasNext()) {
        $element = $iterator->next();
        echo $element . PHP_EOL;
    }
?>