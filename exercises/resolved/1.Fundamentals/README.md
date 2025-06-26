# Resueltos

A continuación se dejan los resueltos de los ejercicios de la sección 1.Fundamentals:

1.
```javascript
db.movies.find({}).pretty();
```

2.
```javascript
db.movies.find({ title: "Back to the Future" }, { directors: 1, _id: 0 });
```

3.
```javascript
db.movies.countDocuments({ title: /Terminator/ });
```

4.
```javascript
db.movies.countDocuments({ runtime: { $gt: 120 } });
```

5.
```javascript
db.movies.find(
  { runtime: { $gt: 200 }, genres: { $all: ["Action", "Sci-Fi"] }, type: 'movie' },
  { year: 1, _id: 0 }
);
```

6.
```javascript
db.movies.find(
  {
    $or: [
      { title: /Batman/ },
      { cast: /Batman/ },
      { fullplot: /Batman/ },
      { genres: /Batman/ }
      ]
  });

db.movies.find({ $text : { $search: "batman" } });
```

7.
```javascript
db.movies.find(
  { $text: { $search: "alien" } },
  { score: { $meta: "textScore" }, title: 1, _id: false }
).sort({ score: { $meta: "textScore" } }).limit(15);
```