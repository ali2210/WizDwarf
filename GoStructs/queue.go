package GoStructs




type QueueList struct{
	ELement string
	Length int
	next string
}

	
type List struct{
	QueueList
	len int
}



func (q QueueList)Enque(ele string)QueueList{
	// Enque Code

	if q.isQueueNil(){
		q.ELement = ele
		q.Length = q.Length + 1
	}else{
		q.ELement = ele
		q.Length = q.Length + 1
	}

	return q
}


func (q QueueList)DeQueue()QueueList{

	if q.ELement != "---"{
		q.next = "nil"
	}
	return q

}

func (q QueueList)Print(){
	println("Data print:", q.ELement)
}

func (q QueueList)isCapacityFull(len int)(bool){

	if q.Length == 	len{
		return true
	}
	return false
}

func (q QueueList)isQueueNil()(bool){

	if q.Length == 0{
		return true
	}
	return false
}


func (l List)Add(q QueueList) List{

	var len int = 0
	mylist := List{}
		if l.len == 0{
		  	l.len +=2
		  	len = 0 
		  	mylist.QueueList = q
		  	mylist.len = len
		}
	len +=1
			mylist.QueueList = q
		  	mylist.len = len
	return mylist
}


func (l List)Delete(q QueueList) *List{

	if l.Find(q) {
		l.QueueList.ELement = ""
		l.QueueList.Length-=1 
	}
	return &l
}

func (l List)Find(q QueueList)bool{
	
	if l.QueueList == q{
		return true
	}
	return false 
}