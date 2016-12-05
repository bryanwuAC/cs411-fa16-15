select 
	DISTINCT paper.title, 
	paper.URL 
from 
	paper, 
	writtenby, 
	favorite, 
	people 
where 
	paper.id = writtenby.paper 
	and writtenby.PERSON = people.ID 
	and people.name in
		(select name from people, paper, writtenby, favorite 
			where favorite.Title = paper.Title
			and paper.id = writtenby.paper 
			and writtenby.PERSON = people.ID)
	and paper.tag in 
		(select tag from people, paper, writtenby, favorite 
			where favorite.Title = paper.Title
			and paper.id = writtenby.paper 
			and writtenby.PERSON = people.ID)
