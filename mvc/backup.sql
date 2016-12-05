select 
	DISTINCT 
	paper.title, 
	paper.URL 
from 
	paper,favorite,writtenby,people
where 
 	favorite.User = "Bryan" 
	and favorite.Title = paper.title 
	and paper.id = writtenby.paper 
	and writtenby.PERSON = people.ID 
	and paper.tag in (
		select 
			tag 
		from 
			paper
		where 
			paper.Title = favorite.Title
	)